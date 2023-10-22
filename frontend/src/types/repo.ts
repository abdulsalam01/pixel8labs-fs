export interface Repo {
  title: string;
  desc?: string;
}

export interface Repository {
  id: number;
  name: string;
  full_name: string;
  private: boolean;
  description: string | null;
  topics: string[];
  updated_at: string; // You can use a Date type if you prefer
  language: string;
  html_url: string;
}
