export const useLogout = () => {
  const doLogout = async () => {
    const token = localStorage.getItem("token");
    const url = `${process.env.NEXT_PUBLIC_BASE_URL}/logout`;

    try {
      await fetch(url, {
        method: "POST",
        headers: {
          Authorization: `Bearer ${token}`,
          "content-type": "application/json",
        },
      });
      localStorage.removeItem("token");
      window.location.href = "/";
    } catch (error: any) {
      throw new Error("Logout error" + error.message);
    }
  };

  return doLogout;
};
