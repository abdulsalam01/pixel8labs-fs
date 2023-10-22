export const getConfig = () => {
  let config: any = {};
  const token = localStorage.getItem("token");

  if (token !== null) {
    config = {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    };
  }

  return config;
};
