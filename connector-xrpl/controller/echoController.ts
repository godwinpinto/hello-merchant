import express, { Express, Request, Response } from 'express';
import dotenv from 'dotenv';
import { GenericResponse } from '../exception/responseJson';
import { PixResponse } from '../ripple/commonInterfaces';
import { admin } from '../utils/firebaseConfig'


const router = express.Router();

router.post('/notification', async (req: Request, res: Response) => {

    const notification_options = {
        priority: "high",
        timeToLive: 60 * 60 * 24
    };

    const { notification_id, title, body, sound } = req.body

    const message_notification = {
        notification: {
            title: title,
            body: body,
            sound: sound
        }, data: {
            title: title,
            body: body,
            sound: sound
        }
    };

    admin.messaging().sendToDevice(notification_id, message_notification, notification_options)
        .then(response => {
            console.log("response", response);
            //       res.status(200).send("Notification sent successfully")

        })
        .catch(error => {
            console.log(error);
        });
    const response = {
        "status": "ok"
    }
    res.json({ response });
});





export default router