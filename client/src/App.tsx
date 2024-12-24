import './App.css'
import { useEffect, useState } from 'react'
import WebApp from '@twa-dev/sdk'
import axios from 'axios'
import { Text } from '@chakra-ui/react'

interface UserData {
  id: number
  first_name: string
  last_name?: string
  user_name?: string
}

export default function App() {
  const [userData, setUserData] = useState<UserData>();
  useEffect(() => {
    if (WebApp.initDataUnsafe.user) {
      setUserData(WebApp.initDataUnsafe.user as UserData)
    }
  }, []);

  useEffect(() => {
    const createUser = async () => {
      const url = "https://02d0-185-78-137-122.ngrok-free.app/api/v1/user";

      try {
        const res = await axios.post(url, userData);
        console.log(res.data);
      } catch (error) {
        console.error(error);
      }
    };

    createUser();
  }, []);

  return (
    <>
      <Text>{ userData?.id }</Text>
      <Text>{ userData?.first_name }</Text>
      <Text>{ userData?.last_name }</Text>
      <Text>{ userData?.user_name }</Text>
    </>
  )
}
