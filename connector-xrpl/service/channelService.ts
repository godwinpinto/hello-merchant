import jsonpath from "jsonpath";
import { sendWebNotification } from "../ripple/sendWebNotification";

export const processNotification = async (event: any) => {
    try {
        console.log("event.data", event.data)
        const jsonData = JSON.parse(event.data);

        const transactionType = jsonpath.query(jsonData, '$.transaction.TransactionType')[0];
        console.log("transactionType", transactionType)
        if (transactionType == "Payment") {
            const recepientAccountNo = jsonpath.query(jsonData, '$.transaction.Destination')[0];
/*             const devices: Array<any> = await fetchAccountsByAccountNo(recepientAccountNo);
            console.log("devices", devices)
            for (const device of devices) {
 */
                sendWebNotification(event.data, "hello.merchant@coauth.dev", JSON.stringify({"sound":true,lang:"en"}));

                 
                sendWebNotification(event.data, "godwin.pinto@cmss.in", JSON.stringify({"sound":true,lang:"en"}));
/*             }
 */        }
    } catch (e) {
        console.log(e)
    }

}


