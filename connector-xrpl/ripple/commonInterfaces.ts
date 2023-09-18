import { Wallet } from "xrpl";

export interface PixResponse{
    xrpl_account_no: string;
    pix_id: string;
    currency_accepted?:string;
    additional_data?:string;
}

export interface IFundResults{
    wallet: Wallet;
    balance: number;
}