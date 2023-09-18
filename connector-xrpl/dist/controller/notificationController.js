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
const notificationService_1 = require("../service/notificationService");
const router = express_1.default.Router();
router.post('/validate-login', (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    const response = {};
    const { origin_id } = req.body;
    try {
        const result = yield (0, notificationService_1.validateLogin)(origin_id);
        response.status = 200;
        response.data = {
            status: result
        };
    }
    catch (e) {
        response.status = 300;
        response.msg = e.message || "Unknown error";
    }
    res.json({ response });
}));
router.post('/register', (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    const response = {};
    const { account_no, origin_id, origin_add_details, contact_id, contact_type, subscription_details, msg_meta_info, sub_expiry } = req.body;
    const registerNotificationInput = {
        account_no, origin_id, origin_add_details, contact_id, contact_type, subscription_details, msg_meta_info, sub_expiry
    };
    try {
        const result = yield (0, notificationService_1.registerChannel)(registerNotificationInput);
        response.status = 200;
        response.data = {
            txn_id: result
        };
    }
    catch (e) {
        response.status = 300;
        response.msg = e.message || "Unknown error";
    }
    res.json({ response });
}));
router.post('/dashboard', (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    const response = {};
    const { origin_id } = req.body;
    try {
        const result = yield (0, notificationService_1.getAccountDetails)(origin_id);
        response.status = 200;
        response.data = {
            result
        };
    }
    catch (e) {
        response.status = 300;
        response.msg = e.message || "Unknown error";
    }
    res.json({ response });
}));
router.post('/verify-account', (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    const response = {};
    const { account_no } = req.body;
    try {
        const result = yield (0, notificationService_1.verifyAccountService)(account_no);
        response.status = 200;
        response.data = {
            result
        };
    }
    catch (e) {
        response.status = 300;
        response.msg = e.message || "Unknown error";
    }
    res.json({ response });
}));
router.post('/delete', (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    const response = {};
    const { origin_id } = req.body;
    try {
        const result = yield (0, notificationService_1.deleteChannel)(origin_id);
        response.status = 200;
        response.data = {
            txn_id: result
        };
    }
    catch (e) {
        response.status = 300;
        response.msg = e.message || "Unknown error";
    }
    res.json({ response });
}));
exports.default = router;
