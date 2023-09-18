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
exports.getLast10Transactions = exports.walletBalance = void 0;
const xrpl_1 = require("xrpl");
const constants_1 = require("../utils/constants");
const walletBalance = (walletAddress) => __awaiter(void 0, void 0, void 0, function* () {
    const client = new xrpl_1.Client("wss://" + constants_1.XRPL_NETWORK);
    let results;
    try {
        yield client.connect();
        console.log("2");
        const command = {
            account: walletAddress,
            command: "account_info",
            ledger_index: "validated"
        };
        const balance_result = yield client.request(command);
        console.log("balance_result", balance_result);
        return balance_result;
    }
    catch (e) {
        console.log(e);
    }
    finally {
        try {
            client.disconnect();
        }
        catch (e) {
        }
    }
    return null;
});
exports.walletBalance = walletBalance;
const getLast10Transactions = (walletAddress) => __awaiter(void 0, void 0, void 0, function* () {
    const client = new xrpl_1.Client("wss://" + constants_1.XRPL_NETWORK);
    let results;
    try {
        yield client.connect();
        console.log("2");
        const command = {
            id: 1,
            account: walletAddress,
            command: "account_tx",
            "ledger_index_min": -1,
            "ledger_index_max": -1,
            "binary": false,
            "limit": 10,
            "forward": false
        };
        const history_result = yield client.request(command);
        console.log("history_result", history_result);
        return history_result;
    }
    catch (e) {
        console.log(e);
    }
    finally {
        try {
            client.disconnect();
        }
        catch (e) {
        }
    }
    return null;
});
exports.getLast10Transactions = getLast10Transactions;
