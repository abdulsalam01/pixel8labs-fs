// utils/auth.js
import axiosInstance from './http'; // Import the Axios instance

export const login = async() => {
  try {
    // Perform the login operation to obtain the token
    const response = await axiosInstance.post('/login');

    // Set the url in the Axios instance's headers
    const url = response.data;

    // Return any additional user data if needed
    console.log(url, "INI")
    return url;
  } catch (error) {
    // Handle login errors
    throw error;
  }
};


/**
 * Helper Authentication of user.
 * User helper
 */
type UserAccount = {
	ID:           String;
	AvatarUrl:    String;
	Name:         String;
	Username:     String;
	Email:        String;
	Bio:          String;
	Follower:     Number;
	FollowersUrl: String;
}

export function setUserDataInLocalStorage(userData: UserAccount): Boolean {
  try {
    const serializedUserData = JSON.stringify(userData);
    window.localStorage.setItem('userData', serializedUserData);

    return true;
  } catch (error) {
    return false;
  }
}

// Get user data from local storage
export function getUserDataFromLocalStorage(): UserAccount {
  try {
    const userData = window.localStorage.getItem('userData');

    return userData ? JSON.parse(userData) : null;
  } catch (error) {
    return {} as UserAccount;
  }
}

// Remove user data from local storage
export function removeUserDataFromLocalStorage(): Boolean {
  try {
    window.localStorage.removeItem('userData');

    return true;
  } catch (error) {
    return false;
  }
}
