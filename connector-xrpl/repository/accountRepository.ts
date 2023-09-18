import { PrismaClient } from '@prisma/client'
const prisma = new PrismaClient()

export const createTransaction = async (record: any): Promise<any> => {
    const result = await prisma.uPN_TRANSACTION_MASTER.create({
        data: record,
    })
    console.log("create", result)
    return result;
}

/* export const fetchXrplAccountByAccountId = async (accountId: string):Promise<any> => {
    const result = await prisma.xRPL_ACCOUNT_MASTER.findFirst({
        where: { XAM_ROW_ID: accountId }
    })
    console.log("fetch", result);
    return result;
} */

