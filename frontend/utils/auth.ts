// utils/auth.js
import axiosInstance from './http'; // Import the Axios instance

export const login = async (username: string, password: string) => {
  try {
    // Perform the login operation to obtain the token
    const response = await axiosInstance.post('/login', { username, password });

    // Set the token in the Axios instance's headers
    const token = response.data.token;
    axiosInstance.defaults.headers.common['Authorization'] = `Bearer ${token}`;

    // Return any additional user data if needed
    return response.data.user;
  } catch (error) {
    // Handle login errors
    throw error;
  }
};
