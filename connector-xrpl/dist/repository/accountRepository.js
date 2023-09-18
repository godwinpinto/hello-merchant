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
exports.createTransaction = void 0;
const client_1 = require("@prisma/client");
const prisma = new client_1.PrismaClient();
const createTransaction = (record) => __awaiter(void 0, void 0, void 0, function* () {
    const result = yield prisma.uPN_TRANSACTION_MASTER.create({
        data: record,
    });
    console.log("create", result);
    return result;
});
exports.createTransaction = createTransaction;
/* export const fetchXrplAccountByAccountId = async (accountId: string):Promise<any> => {
    const result = await prisma.xRPL_ACCOUNT_MASTER.findFirst({
        where: { XAM_ROW_ID: accountId }
    })
    console.log("fetch", result);
    return result;
} */
