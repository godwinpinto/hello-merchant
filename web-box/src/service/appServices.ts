import axios from 'axios';

const SERVER_URL=import.meta.env.VITE_RIPPLE_ECHO_URL;

const headersVal={headers:{"Content-Type":"application/json"}}

export const fetchData = async (email_address:string):Promise<boolean> => {
    try {
        const response = await axios.post(SERVER_URL+"/notification/validate-login",{origin_id:email_address},headersVal);
        if(response.status==200){
            return response.data.response.data.status
        }else{
            return false
        }
    } catch (error) {
        console.error('Error fetching data:', error);
        return false
    }
}

export const registerUser = async (input:any):Promise<boolean> => {
    try {
        const response = await axios.post(SERVER_URL+"/notification/register",input,headersVal);
        if(response.data.response.status==200){
            return true
        }else{
            return false
        }
    } catch (error) {
        console.error('Error fetching data:', error);
        return false
    }
}
export const fetchTransactions = async (origin_id:any):Promise<any> => {
    try {
        const response = await axios.post(SERVER_URL+"/transactions",{origin_id:origin_id},headersVal);
        return response
    } catch (error) {
        console.error('Error fetching data:', error);
        return null
    }
}


export const verifyAccountNo = async (account_no:any):Promise<any> => {
    try {
        const response = await axios.post(SERVER_URL+"/notification/verify-account",{account_no:account_no},headersVal);
        return response
    } catch (error) {
        console.error('Error fetching data:', error);
        return null
    }
}