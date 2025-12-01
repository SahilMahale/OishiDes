import { QueryFunctionContext } from '@tanstack/react-query';
import axios from 'axios';

export interface Bookings {
  bookingID: string
  user: string
}
const getBookingsList = async ({ queryKey }: QueryFunctionContext<[string, string, string]>): Promise<Bookings[]> => {
  const [, authKey, user] = queryKey;
  let API_GATEWAY = process.env.API_GATEWAY
  let API_PORT = process.env.API_PORT
  if (API_GATEWAY === "" || API_GATEWAY === "localhost") {
    API_GATEWAY = "http://localhost"
  }
  if (API_PORT === "") {
    API_PORT = "8080"
  }
  const resp = await axios
    .get(`/api/bookings`, {
      headers: {
        'Content-Type': 'application/type',
        Authorization: 'Bearer ' + authKey,
      },
      params: {
        user,
      },
    })

  return resp.data;
};
const getUsersList = async (): Promise<Array<string>> => {
  const data = await axios
    .get(`/api/user/info`, {
      headers: {
        'Content-Type': 'application/type',
      },
    });

  return data.data;
};

const putBookings = async (user: string, tickets: number, authKey: string) => {
  let data;
  await axios
    .post(
      `$/api/bookings`,
      {
        user: user,
        tickets: tickets,
      },
      {
        headers: {
          'Content-Type': 'application/json',
          Authorization: 'Bearer ' + authKey,
        },
      }
    )
    .then((Response) => {
      data = Response.data;
    });

  return data;
};
const userSignup = async (user: string, email: string, pass: string, isadmin: boolean) => {
  let data;
  await axios
    .post(
      `/api/user/signup`,
      {
        user: user,
        email: email,
        pass: pass,
        isadmin: isadmin,
      },
      {
        headers: {
          'Content-Type': 'application/json',
        },
      }
    )
    .then((Response) => {
      data = Response.data;
    });

  return data;
};
type authObj = {
  auth_token: string
}
const userLogin = async (user: string, pass: string): Promise<authObj> => {
  //let data;
  const resp = await axios
    .post(
      `/api/user/signin`,
      {
        user: user,
        pass: pass,
      },
      {
        headers: {
          'Content-Type': 'application/json',
        },
      }
    )
    /* .then((Response) => {
      // console.log(Response.data)
      data = new Promise()
    }) */;

  return resp.data;
};

export { getBookingsList, getUsersList, putBookings, userLogin, userSignup };
