import React from "react";

const Visitor = () => {
  return (
    <div>
      <h2>Latest Visitors</h2>
      <div className="d-flex">
        <img src="/user1.jpg" alt="User 1" className="img-fluid rounded-circle me-2" />
        <img src="/user2.jpg" alt="User 2" className="img-fluid rounded-circle me-2" />
        <img src="/user3.jpg" alt="User 3" className="img-fluid rounded-circle" />
      </div>
    </div>
  );
}

export default Visitor;
