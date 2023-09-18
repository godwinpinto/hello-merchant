import * as admin from "firebase-admin";
import * as serviceAccount from "../utils/pks.json";
import { firebaseConfig } from "../utils/pks";


admin.initializeApp({
  credential: admin.credential.cert(serviceAccount as admin.ServiceAccount)
});

export {admin}