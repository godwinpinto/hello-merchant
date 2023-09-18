import xrpl, { AccountInfoRequest, AccountTxRequest, Client, Wallet, xrpToDrops } from "xrpl";
import { XRPL_NETWORK } from '../utils/constants'
import { IFundResults } from "./commonInterfaces";

export const walletBalance = async (walletAddress: string): Promise<any | null> => {
    const client = new Client("wss://" + XRPL_NETWORK)
    let results: IFundResults;
    try {
        await client.connect();
        console.log("2")

        const command: AccountInfoRequest = {
            account: walletAddress,
            command: "account_info",
            ledger_index: "validated"
        }
        const balance_result = await client.request(command);

        console.log("balance_result", balance_result);

        return balance_result;
    } catch (e: any) {
        console.log(e)
    } finally {
        try {
            client.disconnect()
        } catch (e) {
        }
    }
    return null;

}

export const getLast10Transactions = async (walletAddress: string): Promise<any | null> => {
    const client = new Client("wss://" + XRPL_NETWORK)
    let results: IFundResults;
    try {
        await client.connect();
        console.log("2")

        const command: AccountTxRequest = {
            id:1,
            account: walletAddress,
            command: "account_tx",
            "ledger_index_min": -1,
            "ledger_index_max": -1,
            "binary": false,
            "limit": 10,
            "forward": false
        }
        const history_result = await client.request(command);

        console.log("history_result", history_result);

        return history_result;
    } catch (e: any) {
        console.log(e)
    } finally {
        try {
            client.disconnect()
        } catch (e) {
        }
    }
    return null;

}

