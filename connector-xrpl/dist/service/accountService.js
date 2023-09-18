"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.sendMoney = exports.getBalance = exports.authenticateUserByPixId = exports.registerUser = void 0;
const accountRepository_1 = require("../repository/accountRepository");
const rpayRepository_1 = require("../repository/rpayRepository");
const snowflake_uuid_1 = require("snowflake-uuid");
const sha512_crypt_ts_1 = require("sha512-crypt-ts");
const xrpl_1 = require("xrpl");
const client_1 = require("@prisma/client");
const walletUtils_1 = require("../ripple/walletUtils");
const prisma = new client_1.PrismaClient();
const generator = new snowflake_uuid_1.Worker(0, 1, {
    workerIdBits: 5,
    datacenterIdBits: 5,
    sequenceBits: 12,
});
const registerUser = (identifier, pin) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        const pixData = yield (0, rpayRepository_1.fetchRpayByPixId)(identifier);
        if (!pixData || pixData == null) {
            const wallet = xrpl_1.Wallet.generate();
            console.log("wallet", wallet);
            const currentDate = new Date();
            const rowId = generator.nextId().toString();
            const xrplAccountEntity = {
                XAM_ROW_ID: rowId,
                AC_NO: wallet.address,
                SEED: wallet.seed,
                ACTIVE: "Y",
                CREATED_DT: currentDate.toISOString(),
                CREATED_BY: "SYSTEM",
                UPDATED_DT: currentDate.toISOString(),
                UPDATED_BY: "SYSTEM"
            };
            const rPayAccountEntity = {
                XPAM_ROW_ID: rowId,
                XAM_ROW_ID: rowId,
                PIX_ID: identifier,
                PIN: sha512_crypt_ts_1.sha512.crypt(pin, rowId.substring(0, 8)),
                ATTEMPTS: 0,
                ACTIVE: "Y",
                CREATED_DT: currentDate.toISOString(),
                CREATED_BY: "SYSTEM",
                UPDATE_DT: currentDate.toISOString(),
                UPDATE_BY: "SYSTEM"
            };
            const transactionStatus = yield prisma.$transaction((tx) => __awaiter(void 0, void 0, void 0, function* () {
                yield (0, accountRepository_1.createXrplAccount)(xrplAccountEntity, tx);
                yield (0, rpayRepository_1.createPixId)(rPayAccountEntity, tx);
                return true;
            }));
            if (transactionStatus) {
                console.log("Transaction success");
                const results = yield (0, walletUtils_1.fundWallet)(wallet.seed || '');
                return results;
            }
            else {
                throw new Error("Transaction Failed");
            }
        }
        else {
            throw new Error("Account already Registered");
        }
    }
    catch (e) {
        throw new Error(e.message || "Unknown error");
    }
});
exports.registerUser = registerUser;
const authenticateUserByPixId = (identifier, pin) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        const pixData = yield (0, rpayRepository_1.fetchRpayByPixId)(identifier);
        if (pixData) {
            if (!pin && pixData.PIN != sha512_crypt_ts_1.sha512.crypt(pin, pixData.XPAM_ROW_ID.substring(0, 8))) {
                return false;
            }
            else {
                return true;
            }
        }
        else {
            return false;
        }
    }
    catch (e) {
        throw new Error(e.message || "Unknown error");
    }
});
exports.authenticateUserByPixId = authenticateUserByPixId;
const getBalance = (identifier) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        const pixData = yield (0, rpayRepository_1.fetchRpayByPixId)(identifier);
        if (pixData) {
            const acData = yield (0, accountRepository_1.fetchXrplAccountByAccountId)(pixData.XAM_ROW_ID);
            if (acData) {
                const acBalance = yield (0, walletUtils_1.walletBalance)(acData.AC_NO);
                return acBalance;
            }
            else {
                throw new Error("Invalid Account");
            }
        }
        else {
            throw new Error("Invalid Account");
        }
    }
    catch (e) {
        throw new Error(e.message || "Unknown error");
    }
});
exports.getBalance = getBalance;
const sendMoney = (identifier, pixId, amount) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        const pixData = yield (0, rpayRepository_1.fetchRpayByPixId)(identifier);
        if (pixData) {
            const acData = yield (0, accountRepository_1.fetchXrplAccountByAccountId)(pixData.XAM_ROW_ID);
            if (acData) {
                //const acBalance=await walletBalance(acData.AC_NO)
                const paymentStatus = (0, walletUtils_1.makePayment)(acData.SEED, "rPTPn2Pxx824YVzcjHvUeXx7V3wkCFP52a", parseFloat(amount));
                return paymentStatus;
            }
            else {
                throw new Error("Invalid Account");
            }
        }
        else {
            throw new Error("Invalid Account");
        }
    }
    catch (e) {
        throw new Error(e.message || "Unknown error");
    }
});
exports.sendMoney = sendMoney;
