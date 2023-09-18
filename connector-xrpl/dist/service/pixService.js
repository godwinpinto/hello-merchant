"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const client_1 = require("@prisma/client");
const prisma = new client_1.PrismaClient();
/* export const getPixDetails = async (pixId: string): Promise<PixResponse> => {

    try {
        const pixData = await fetchRpayByPixId(pixId);
        if (pixData) {
            const acData = await fetchXrplAccountByAccountId(pixData.XAM_ROW_ID)
            if (acData) {
                const pixResponse: PixResponse = {
                    pix_id: pixId,
                    xrpl_account_no: acData.AC_NO
                }
                return pixResponse;
            } else {
                throw new Error("Invalid Account")
            }
        } else {
            throw new Error("Invalid Account")
        }
    } catch (e: any) {
        throw new Error(e.message || "Unknown error")
    }
}

 */ 
