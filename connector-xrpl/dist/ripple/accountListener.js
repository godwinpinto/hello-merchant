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
exports.doSubscribe = exports.startListeningAccounts = void 0;
const constants_1 = require("../utils/constants");
const ws_1 = __importDefault(require("ws"));
const channelService_1 = require("../service/channelService");
const accountRepository_1 = require("../repository/accountRepository");
let socket;
const AWAITING = {};
let autoid_n = 0;
const startListeningAccounts = () => __awaiter(void 0, void 0, void 0, function* () {
    socket = new ws_1.default('wss://s.altnet.rippletest.net:51233');
    const subscribers = yield (0, accountRepository_1.fetchDistinctAccounts)();
    const accountArray = subscribers.map((item) => item.XRPL_AC_NO).filter((account) => account.startsWith('r'));
    //    const accountArray=["rDKH6NniQpqoAJNBh4bTf7y9rXigwkyZHa"]
    console.log(accountArray);
    console.log(constants_1.XRPL_NETWORK);
    socket.addEventListener('close', (event) => {
        console.log('Disconnected...');
        (0, exports.startListeningAccounts)();
    });
    const handleResponse = (data) => {
        if (!data.hasOwnProperty("id")) {
            console.error("Got response event without ID:", data);
            return;
        }
        if (AWAITING.hasOwnProperty(data.id)) {
            AWAITING[data.id].resolve(data);
        }
        else {
            console.warn("Response to un-awaited request w/ ID " + data.id);
        }
    };
    const WS_HANDLERS = {
        "response": handleResponse
    };
    socket.addEventListener('message', (event) => {
        console.log('Got message from server:', event.data);
        (0, channelService_1.processNotification)(event);
        const parsed_data = JSON.parse(event.data);
        if (WS_HANDLERS.hasOwnProperty(parsed_data.type)) {
            WS_HANDLERS[parsed_data.type](parsed_data);
        }
        else {
            console.log("Unhandled message from server", event);
        }
    });
    WS_HANDLERS["transaction"] = log_tx;
    socket.addEventListener('open', (event) => {
        // This callback runs when the connection is open
        console.log("Connected!");
        pingpong();
        (0, exports.doSubscribe)(accountArray);
    });
});
exports.startListeningAccounts = startListeningAccounts;
const doSubscribe = (accountArray) => __awaiter(void 0, void 0, void 0, function* () {
    const sub_response = yield api_request({
        command: "subscribe",
        accounts: accountArray
    });
    /*     if (sub_response.status === "success") {
            console.log("Successfully subscribed!")
        } else {
            console.error("Error subscribing: ", sub_response)
        } */
});
exports.doSubscribe = doSubscribe;
const log_tx = function (tx) {
    console.log(tx.transaction.TransactionType + " transaction sent by " +
        tx.transaction.Account +
        "\n  Result: " + tx.meta.TransactionResult +
        " in ledger " + tx.ledger_index +
        "\n  Validated? " + tx.validated);
};
const api_request = (options) => __awaiter(void 0, void 0, void 0, function* () {
    if (socket.readyState === 0) {
        console.error("Socket is not connected yet");
        return "issue";
    }
    try {
        socket.send(JSON.stringify(options));
    }
    catch (error) {
        //reject(error)
    }
});
const pingpong = () => __awaiter(void 0, void 0, void 0, function* () {
    console.log("Ping...");
    const response = yield api_request({ command: "ping" });
    console.log("Pong!", response);
});
const addNewAccount = () => __awaiter(void 0, void 0, void 0, function* () {
    const subscribers = yield (0, accountRepository_1.fetchDistinctAccounts)();
    const accountArray = subscribers.map((item) => item.XRPL_AC_NO).filter((account) => account.startsWith('r'));
    (0, exports.doSubscribe)(accountArray);
});
setInterval(function () {
    addNewAccount();
}, 15 * 1000);
/* setInterval(function () {
    pingpong()
    if (socket.readyState == socket.OPEN) {
        socket.close();
    }
}, 60 * 1000);
 */
