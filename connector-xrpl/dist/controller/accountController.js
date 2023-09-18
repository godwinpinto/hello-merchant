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
const accountService_1 = require("../service/accountService");
const router = express_1.default.Router();
router.post('/register', (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    const response = {};
    try {
        const { identifier, pin } = req.body;
        console.log("identifier", identifier, pin);
        if (!identifier || !pin || identifier.trim() == "" || pin.trim() == "") {
            throw new Error("Invalid request data");
        }
        const results = yield (0, accountService_1.registerUser)(identifier.trim(), pin.trim());
        response.status = 200;
        response.data = results;
    }
    catch (e) {
        console.log("error", e);
        response.status = 500;
        response.msg = e.message || "An error occurred";
    }
    res.json({ response });
}));
router.post('/balance', (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    const response = {};
    try {
        const { identifier, pin } = req.body;
        console.log("identifier", identifier, pin);
        if (!identifier || !pin || identifier.trim() == "" || pin.trim() == "") {
            throw new Error("Invalid request data");
        }
        if (!(yield (0, accountService_1.authenticateUserByPixId)(identifier, pin))) {
            throw new Error("Invalid login / pin");
        }
        const results = yield (0, accountService_1.getBalance)(identifier.trim());
        response.status = 200;
        response.data = results;
    }
    catch (e) {
        console.log("error", e);
        response.status = 500;
        response.msg = e.message || "An error occurred";
    }
    res.json({ response });
}));
router.post('/transfer', (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    const response = {};
    try {
        const { identifier, pin, destinationPixId, amount } = req.body;
        console.log("identifier", identifier, pin, destinationPixId, amount);
        if (!identifier || !pin || !destinationPixId || !amount || identifier.trim() == "" || pin.trim() == "" || destinationPixId.trim() == "" || amount.trim() == "") {
            throw new Error("Invalid request data");
        }
        if (!(yield (0, accountService_1.authenticateUserByPixId)(identifier, pin))) {
            throw new Error("Invalid login / pin");
        }
        const results = yield (0, accountService_1.sendMoney)(identifier.trim(), destinationPixId.trim(), amount.trim());
        response.status = 200;
        response.data = results;
    }
    catch (e) {
        console.log("error", e);
        response.status = 500;
        response.msg = e.message || "An error occurred";
    }
    res.json({ response });
}));
exports.default = router;
