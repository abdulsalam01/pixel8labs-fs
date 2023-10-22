import { AppContext, type AppContextType } from "@/context/AppContext";
import { getConfig } from "@/helper/getConfig";
import { type Repository } from "@/types";
import { useContext, useEffect } from "react";

export const useGetRepo = () => {
  const ctx = useContext(AppContext) as AppContextType;

  useEffect(() => {
    const config = getConfig();
    const url = `${process.env.NEXT_PUBLIC_BASE_URL}/repo`;
    fetch(url, config)
      .then(async (res) => {
        return await res.json();
      })
      .then((data) => {
        const dataOfUser = data.data as Repository[];
        ctx.setRepos(dataOfUser);
      })
      .catch((err) => {
        throw err;
      });
  }, []);
};
