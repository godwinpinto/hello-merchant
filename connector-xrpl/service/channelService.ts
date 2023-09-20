import jsonpath from "jsonpath";
import { sendWebNotification } from "../ripple/sendWebNotification";
import { UPN_TRANSACTION_LOG } from "@prisma/client";
import { Worker } from 'snowflake-uuid';
import { createTransactionLog, fetchAccountsByAccountNo } from "../repository/accountRepository";
import axios from "axios";


const generator = new Worker(0, 1, {
    workerIdBits: 5,
    datacenterIdBits: 5,
    sequenceBits: 12,
});



export const processNotification = async (event: any) => {
    try {
        console.log("event.data", event.data)
        const jsonData = JSON.parse(event.data);

        const transactionType = jsonpath.query(jsonData, '$.transaction.TransactionType')[0];
        console.log("transactionType", transactionType)
        if (transactionType == "Payment") {
            const recepientAccountNo = jsonpath.query(jsonData, '$.transaction.Destination')[0];
            const devices: Array<any> = await fetchAccountsByAccountNo(recepientAccountNo);
            console.log("devices", devices)
            for (const device of devices) {
                const uniqueId = generator.nextId().toString();
                const redactData = await redactService(jsonData);
                if (redactData != "") {
                    const dbObject: UPN_TRANSACTION_LOG = {
                        UTL_ROW_ID: uniqueId,
                        UUM_ROW_ID: device.UUM_ROW_ID,
                        LOG_DATA: redactData
                    };
                    createTransactionLog(dbObject)
                }
                console.log("device.UUM_ROW_ID",device.UUM_ROW_ID)
                sendWebNotification(event.data, device.UUM_ROW_ID, JSON.stringify({ "sound": true, lang: "en" }), uniqueId);
            }
        }
    } catch (e) {
        console.log(e)
    }

}


const redactService = async (log: JSON): Promise<string> => {
    try {
        const response = await axios.post(process.env.HELLO_ADMIN_URL + "/api/redact", { body: JSON.stringify(log) });
        if (response.status == 200) {
            return JSON.stringify(response.data);
        }
        return ""
    } catch (error) {
        console.error('Error fetching data:', error);
        return ""
    }

} 