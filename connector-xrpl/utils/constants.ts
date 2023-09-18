
export const XRPL_NETWORK="s.altnet.rippletest.net:51233"
//export const XRPL_NETWORK="s.altnet.rippletest.net 51235"

export interface RegisterNotificationInput{
    account_no: string
    origin_id: string
    origin_add_details:string
    contact_id:string
    contact_type:string
    subscription_details:string
    msg_meta_info:string
    sub_expiry: Date
}

export const addMonths=(date:Date, months:number):Date => {
    date.setMonth(date.getMonth() + months);
  
    return date;
  }
