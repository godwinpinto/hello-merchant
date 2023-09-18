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
const express_1 = __importDefault(require("express"));
const firebaseConfig_1 = require("../utils/firebaseConfig");
const router = express_1.default.Router();
router.post('/notification', (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    const notification_options = {
        priority: "high",
        timeToLive: 60 * 60 * 24
    };
    const { notification_id, title, body, sound } = req.body;
    const message_notification = {
        notification: {
            title: title,
            body: body,
            sound: sound
        }, data: {
            title: title,
            body: body,
            sound: sound
        }
    };
    firebaseConfig_1.admin.messaging().sendToDevice(notification_id, message_notification, notification_options)
        .then(response => {
        console.log("response", response);
        //       res.status(200).send("Notification sent successfully")
    })
        .catch(error => {
        console.log(error);
    });
    const response = {
        "status": "ok"
    };
    res.json({ response });
}));
exports.default = router;
