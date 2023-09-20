import { PrismaClient } from '@prisma/client'
const prisma = new PrismaClient()

export const createTransaction = async (record: any): Promise<any> => {
    const result = await prisma.uPN_TRANSACTION_MASTER.create({
        data: record,
    })
    console.log("create", result)
    return result;
}

export const createTransactionLog = async (record: any): Promise<any> => {
    const result = await prisma.uPN_TRANSACTION_LOG.create({
        data: record,
    })
    console.log("create", result)
    return result;
}

export const fetchAccountsByAccountNo = async (account_number: string): Promise<any> => {
    const result = await prisma.uPN_XRPL_MAPPING.findMany({
        where: {
            XRPL_AC_NO: account_number,
            ACTIVE: "Y"
        }
    })
    console.log("fetch", result);
    return result;
}

export const fetchDistinctAccounts = async (): Promise<any> => {
    const result = await prisma.uPN_XRPL_MAPPING.findMany({
        where: { ACTIVE: "Y" },
        distinct: 'XRPL_AC_NO',
        select: { XRPL_AC_NO: true }
    })
    return result;
}

/* export const fetchXrplAccountByAccountId = async (accountId: string):Promise<any> => {
    const result = await prisma.xRPL_ACCOUNT_MASTER.findFirst({
        where: { XAM_ROW_ID: accountId }
    })
    console.log("fetch", result);
    return result;
} */

