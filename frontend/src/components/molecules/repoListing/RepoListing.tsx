import { BadgeSm } from "@/components/atoms";
import { Dot } from "@/components/atoms/badge/Dot";
import { languageColors } from "@/constant/lang";
import { monthNames } from "@/constant/month";
import { type Repository } from "@/types";
import Link from "next/link";

export const RepoListing = ({
  name,
  description,
  language,
  updated_at: updatedAt,
  private: isPrivate,
  html_url: htmlURL,
}: Repository) => {
  const date = new Date(updatedAt);
  const month = monthNames[date.getMonth()];
  const day = date.getDate();
  const formattedDate = `updated on ${month} ${day}`;

  return (
    <Link href={htmlURL}>
      <div className="bg-gray-25 border p-6 rounded-lg mb-6">
        <div className="mb-2 flex items-center gap-3 flex-wrap">
          <h3 className="font-bold text-gray-800">{name}</h3>
          {!isPrivate ? (
            <BadgeSm
              text="public"
              className="bg-indigo-50 text-badge-text-public font-medium"
            />
          ) : (
            <BadgeSm
              text="private"
              className="bg-private text-gray-700 font-medium"
            />
          )}
        </div>
        {/* Description */}
        <p className="text-gray-800 text-sm">{description}</p>

        {/* Repo info */}
        <div className="mt-6 flex items-center gap-6 text-xs">
          {language !== "" ? (
            <div className="flex items-center gap-1">
              <Dot
                style={{ backgroundColor: languageColors[language] }}
                className="mb-[1.2px]"
              />
              <span className="text-gray-800">{language}</span>
            </div>
          ) : null}
          <div className="text-gray-500">{formattedDate}</div>
        </div>
      </div>
    </Link>
  );
};
