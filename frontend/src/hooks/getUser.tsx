import { AppContext, type AppContextType } from "@/context/AppContext";
import { getConfig } from "@/helper/getConfig";
import { type UserData } from "@/types";
import { useContext, useEffect } from "react";

export const useGetUser = () => {
  const ctx = useContext(AppContext) as AppContextType;

  useEffect(() => {
    const config = getConfig();
    const url = `${process.env.NEXT_PUBLIC_BASE_URL}/user`;
    fetch(url, config)
      .then(async (res) => {
        return await res.json();
      })
      .then((data) => {
        const dataOfUser = data.data as UserData;
        ctx.setUserdata(dataOfUser);
      })
      .catch((err) => {
        throw err;
      });
  }, []);
};
