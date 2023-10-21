// components/Profile.js
import React from 'react';
import Visitor from './Visitor';

const Profile = () => {
  return (
    <>
      {/* Profile Section */}
      <div className="d-flex flex-column align-items-center">
        <img src="/user-profile-image.jpg" alt="Profile" className="img-fluid rounded-circle mb-3" />
        <h1 className="h2">John Doe</h1>
        <p className="text-muted">@johndoe</p>
      </div>
      <div className='d-flex flex-column align-items-left'>
        <div className="mt-4">
          <h2>About</h2>
          <p>The about section goes here.</p>
        </div>
        <div className="mt-4">
          <img src="/email-icon.png" alt="Email" className="me-2" />
          <p>johndoe@example.com</p>
          <p>Visitor Information: Some details here</p>
        </div>
        <div className="mt-4">
          <Visitor />
        </div>
      </div>
    </>
  );
};

export default Profile;
