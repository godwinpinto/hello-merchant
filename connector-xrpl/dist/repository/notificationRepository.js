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
exports.fetchAccountsByAccountNo = exports.fetchDistinctAccounts = void 0;
const client_1 = require("@prisma/client");
const prisma = new client_1.PrismaClient();
const fetchDistinctAccounts = () => __awaiter(void 0, void 0, void 0, function* () {
    const result = yield prisma.uPN_XRPL_MAPPING.findMany({
        where: { ACTIVE: "Y" },
        distinct: 'XRPL_AC_NO',
        select: { XRPL_AC_NO: true }
    });
    return result;
});
exports.fetchDistinctAccounts = fetchDistinctAccounts;
/* export const fetchAccountsByDeviceId = async (origin_id: string): Promise<any> => {
    const result = await prisma.uPN_XRPL_MAPPING.findFirst({
        where: { ORIGIN_ID: origin_id, ACTIVE: "Y" },
        select: { XRPL_AC_NO: true,XUCM_ROW_ID:true }
    })
    console.log("fetch", result);
    return result;
} */
const fetchAccountsByAccountNo = (account_number) => __awaiter(void 0, void 0, void 0, function* () {
    const result = yield prisma.uPN_XRPL_MAPPING.findMany({
        where: {
            XRPL_AC_NO: account_number,
            ACTIVE: "Y"
        }
    });
    console.log("fetch", result);
    return result;
});
exports.fetchAccountsByAccountNo = fetchAccountsByAccountNo;
