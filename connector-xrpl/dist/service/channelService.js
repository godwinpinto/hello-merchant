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
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.processNotification = void 0;
const jsonpath_1 = __importDefault(require("jsonpath"));
const sendWebNotification_1 = require("../ripple/sendWebNotification");
const snowflake_uuid_1 = require("snowflake-uuid");
const accountRepository_1 = require("../repository/accountRepository");
const axios_1 = __importDefault(require("axios"));
const generator = new snowflake_uuid_1.Worker(0, 1, {
    workerIdBits: 5,
    datacenterIdBits: 5,
    sequenceBits: 12,
});
const processNotification = (event) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        console.log("event.data", event.data);
        const jsonData = JSON.parse(event.data);
        const transactionType = jsonpath_1.default.query(jsonData, '$.transaction.TransactionType')[0];
        console.log("transactionType", transactionType);
        if (transactionType == "Payment") {
            const recepientAccountNo = jsonpath_1.default.query(jsonData, '$.transaction.Destination')[0];
            const devices = yield (0, accountRepository_1.fetchAccountsByAccountNo)(recepientAccountNo);
            console.log("devices", devices);
            for (const device of devices) {
                const uniqueId = generator.nextId().toString();
                const redactData = yield redactService(jsonData);
                if (redactData != "") {
                    const dbObject = {
                        UTL_ROW_ID: uniqueId,
                        UUM_ROW_ID: device.UUM_ROW_ID,
                        LOG_DATA: redactData
                    };
                    (0, accountRepository_1.createTransactionLog)(dbObject);
                }
                console.log("device.UUM_ROW_ID", device.UUM_ROW_ID);
                (0, sendWebNotification_1.sendWebNotification)(event.data, device.UUM_ROW_ID, JSON.stringify({ "sound": true, lang: "en" }), uniqueId);
            }
        }
    }
    catch (e) {
        console.log(e);
    }
});
exports.processNotification = processNotification;
const redactService = (log) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        const response = yield axios_1.default.post(process.env.HELLO_ADMIN_URL + "/api/redact", { body: JSON.stringify(log) });
        if (response.status == 200) {
            return JSON.stringify(response.data);
        }
        return "";
    }
    catch (error) {
        console.error('Error fetching data:', error);
        return "";
    }
});
