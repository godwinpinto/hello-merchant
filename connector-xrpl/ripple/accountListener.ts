import xrpl, { AccountInfoRequest, Client, Wallet, xrpToDrops } from "xrpl";
import { XRPL_NETWORK } from '../utils/constants'
import WebSocket from 'ws'
import { processNotification } from "../service/channelService";
import { fetchDistinctAccounts } from "../repository/accountRepository";

let socket: WebSocket;

const AWAITING: any = {}

let autoid_n = 0

export const startListeningAccounts = async () => {

    socket = new WebSocket('wss://s.altnet.rippletest.net:51233')

    const subscribers = await fetchDistinctAccounts();
    const accountArray = subscribers.map((item: any) => item.XRPL_AC_NO).filter((account: any) => account.startsWith('r'));
//    const accountArray=["rDKH6NniQpqoAJNBh4bTf7y9rXigwkyZHa"]
    console.log(accountArray)
    console.log(XRPL_NETWORK);
    socket.addEventListener('close', (event: any) => {
        console.log('Disconnected...')
        startListeningAccounts();
    })

    const handleResponse = (data: any) => {
        if (!data.hasOwnProperty("id")) {
            console.error("Got response event without ID:", data)
            return
        }
        if (AWAITING.hasOwnProperty(data.id)) {
            AWAITING[data.id].resolve(data)
        } else {
            console.warn("Response to un-awaited request w/ ID " + data.id)
        }
    }

    const WS_HANDLERS: any = {
        "response": handleResponse
    }
    socket.addEventListener('message', (event: any) => {
        console.log('Got message from server:', event.data)
        processNotification(event);
        const parsed_data = JSON.parse(event.data)
        if (WS_HANDLERS.hasOwnProperty(parsed_data.type)) {
            WS_HANDLERS[parsed_data.type](parsed_data)
        } else {
            console.log("Unhandled message from server", event)
        }
    })



    WS_HANDLERS["transaction"] = log_tx

    socket.addEventListener('open', (event: any) => {
        // This callback runs when the connection is open
        console.log("Connected!")
        pingpong()
        doSubscribe(accountArray)
    })

}

export const doSubscribe = async (accountArray: Array<string>) => {
    const sub_response = await api_request({
        command: "subscribe",
        accounts: accountArray
    })
/*     if (sub_response.status === "success") {
        console.log("Successfully subscribed!")
    } else {
        console.error("Error subscribing: ", sub_response)
    } */
}


const log_tx = function (tx: any) {
    console.log(tx.transaction.TransactionType + " transaction sent by " +
        tx.transaction.Account +
        "\n  Result: " + tx.meta.TransactionResult +
        " in ledger " + tx.ledger_index +
        "\n  Validated? " + tx.validated)
}
const api_request = async (options: any): Promise<any> => {
    if (socket.readyState === 0) {
        console.error("Socket is not connected yet")
        return "issue";
    }
    try {
        socket.send(JSON.stringify(options))
    } catch (error) {
        //reject(error)
    }
}

const pingpong = async () => {
    console.log("Ping...")
    const response = await api_request({ command: "ping" })
    console.log("Pong!", response)
}

const addNewAccount=async()=>{
    const subscribers = await fetchDistinctAccounts();
    const accountArray = subscribers.map((item: any) => item.XRPL_AC_NO).filter((account: any) => account.startsWith('r'));
    doSubscribe(accountArray);
}

setInterval(function () {
    addNewAccount();
}, 15 * 1000);

/* setInterval(function () {
    pingpong()
    if (socket.readyState == socket.OPEN) {
        socket.close();
    }
}, 60 * 1000);
 */

