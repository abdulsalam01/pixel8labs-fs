import { AppContext, type AppContextType } from "@/context/AppContext";
import { getConfig } from "@/helper/getConfig";
import { type Follower } from "@/types";
import { useContext, useEffect } from "react";

export const useGetFollower = () => {
  const ctx = useContext(AppContext) as AppContextType;

  useEffect(() => {
    const config = getConfig();
    if (ctx.userData !== undefined) {
      fetch(ctx.userData.followers_url, config)
        .then(async (res) => {
          return await res.json();
        })
        .then((data: Follower[]) => {
          ctx.setFollower(data);
        })
        .catch((err) => {
          throw err;
        });
    }
  }, [ctx.userData]);
};
