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
exports.sendFcmNotification = void 0;
const jsonpath_1 = __importDefault(require("jsonpath"));
const xrpl_1 = require("xrpl");
const firebaseConfig_1 = require("../utils/firebaseConfig");
const sendFcmNotification = (eventData, deviceDetails) => __awaiter(void 0, void 0, void 0, function* () {
    const jsonData = JSON.parse(eventData);
    const transactionResult = jsonpath_1.default.query(jsonData, '$.engine_result')[0];
    const type = jsonpath_1.default.query(jsonData, '$.type')[0];
    const validated = jsonpath_1.default.query(jsonData, '$.validated')[0];
    const balance = jsonpath_1.default.query(jsonData, '$.meta.AffectedNodes[0].ModifiedNode.FinalFields.Balance')[0];
    const transactionAmount = jsonpath_1.default.query(jsonData, '$.transaction.Amount')[0];
    if (validated == true && transactionResult == "tesSUCCESS" && type == "transaction") {
        const messageText = "You received " + (0, xrpl_1.dropsToXrp)(transactionAmount) + " XRP on RippleEcho";
        console.log(process.env.ACCOUNT_NO, process.env.PUSHER_CHANNEL);
        console.log(messageText);
        const notification_options = {
            priority: "high",
            timeToLive: 60 * 60 * 24
        };
        const message_notification = {
            notification: {
                title: "Ripple Echo: Payment Received",
                body: messageText,
                sound: "english"
            }, data: {
                title: "Ripple Echo: Payment Received",
                body: messageText,
                sound: "english"
            }
        };
        firebaseConfig_1.admin.messaging().sendToDevice(deviceDetails.CONTACT_ID, message_notification, notification_options)
            .then(response => {
            console.log("response", response);
            //       res.status(200).send("Notification sent successfully")
        })
            .catch(error => {
            console.log(error);
        });
    }
    else {
        console.log("something was not right");
    }
});
exports.sendFcmNotification = sendFcmNotification;
