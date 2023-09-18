import express, { Express, Request, Response } from 'express';
import dotenv from 'dotenv';
import pixController from './controller/echoController';
import { startListeningAccounts } from './ripple/accountListener';
import cors from 'cors'

dotenv.config();

const app: Express = express();
const port = process.env.PORT;

app.use(cors())
app.use(express.json());
app.use(express.urlencoded({ extended: false }));

app.use('/send', pixController);


//app.use('/notification', notificationController);


app.get('/', (req: Request, res: Response) => {
  res.send('Ripple Echo Box API is running');
});


/* setInterval(() => {
  // Your background process logic here
  console.log('Background process is running...');
  startListeningAccounts();
}, 60000); // 5000 milliseconds = 5 seconds
 */
  startListeningAccounts();

app.listen(port, () => {
  console.log(`⚡️[server]: Server is running at http://localhost:${port}`);
});
