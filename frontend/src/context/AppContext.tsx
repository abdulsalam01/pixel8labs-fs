import { type Repository, type Follower, type UserData } from "@/types";
import React, { useState, type ReactNode } from "react";

export const AppContext = React.createContext({});

interface Props {
  children?: ReactNode;
}

export interface AppContextType {
  userData?: UserData;
  setUserdata: React.Dispatch<React.SetStateAction<UserData | undefined>>;
  followers?: Follower[];
  setFollower: React.Dispatch<React.SetStateAction<Follower[] | undefined>>;
  repos?: Repository[];
  setRepos: React.Dispatch<React.SetStateAction<Repository[] | undefined>>;
}

export const AppProvider = ({ children }: Props) => {
  const [userData, setUserdata] = useState<UserData>();
  const [followers, setFollower] = useState<Follower[]>();
  const [repos, setRepos] = useState<Repository[]>();

  const data: AppContextType = {
    userData,
    setUserdata,
    followers,
    setFollower,
    repos,
    setRepos,
  };

  return <AppContext.Provider value={data}>{children}</AppContext.Provider>;
};
