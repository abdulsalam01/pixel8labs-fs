import { type UserData } from "@/types";
import Image from "next/image";

interface Props {
  userData?: UserData;
}

export const Profile = ({ userData }: Props) => {
  const avatarUrl = userData?.avatar_url;

  return (
    <div className="flex lg:flex-col items-center">
      <Image
        className="hidden rounded-full lg:block"
        src={`${avatarUrl ?? "/assets/images/profile.png"}`}
        width={160}
        height={160}
        alt="Profile photo"
      />
      <Image
        className="lg:hidden rounded-full"
        src={`${avatarUrl ?? "/assets/images/profile.png"}`}
        width={80}
        height={80}
        alt="Profile photo"
      />
      <div className="ml-4 lg:ml-0 lg:text-center">
        <h2 className="mt-2 text-2xl text-gray-800 font-bold">
          {userData?.name}
        </h2>
        <div className="text-gray-500 -mt-0.5">@{userData?.login}</div>
      </div>
    </div>
  );
};
