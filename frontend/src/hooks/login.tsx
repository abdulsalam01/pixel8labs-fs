export const useLogin = () => {
  const doLogin = async () => {
    try {
      const response = await fetch("http://localhost:8000/api/login", {
        method: "POST",
        headers: {
          "content-type": "application/json",
        },
      });
      const jsonData = await response.json();
      window.location.href = jsonData.data;
    } catch (error) {
      console.error(error);
      throw error;
    }
  };

  return doLogin;
};
