import React from "react";

// components/Footer.js
const Footer = () => {
  return (
    <footer className="footer text-center py-3 bg-light">
      &copy; {new Date().getFullYear()} Pixel8Labs. All rights reserved.
    </footer>
  );
};

export default Footer;
