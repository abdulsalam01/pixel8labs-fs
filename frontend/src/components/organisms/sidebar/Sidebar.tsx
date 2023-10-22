import { About, Profile } from "@/components/molecules";
import { VisitorList } from "./VisitorList";
import { useContext } from "react";
import { AppContext, type AppContextType } from "@/context/AppContext";

export const Sidebar = () => {
  const ctx = useContext(AppContext) as AppContextType;

  return (
    <div className="px-4  flex flex-col justify-center">
      <Profile userData={ctx.userData} />
      <About userData={ctx.userData} />
      <VisitorList />
    </div>
  );
};
