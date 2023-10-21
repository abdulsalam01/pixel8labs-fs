// components/ProjectList.js
import React from 'react';

const Repository = () => {
  const projects = [
    {
      title: 'Sample Project 1',
      isPublic: true,
      description: 'This is a sample project description.',
      language: 'Ruby',
      lastUpdated: 'October 12, 2023',
    },
    {
      title: 'Sample Project 2',
      isPublic: false,
      description: 'Another project description here.',
      language: 'Python',
      lastUpdated: 'October 15, 2023',
    },
    // Add more projects as needed
  ];

  return (
    <>
      <div className='container bg-white p-3'>
        <h1 className="h2 mb-3">John Doe</h1>

        {projects.map((project, index) => (
          <div className="card bg-light mb-3 bordeer-0" key={index}>
            <div className="card-body">
              <h5 className="card-title">
                {project.title}{' '}
                <span className={`badge ${project.isPublic ? 'bg-success' : 'bg-secondary'}`}>
                  {project.isPublic ? 'Public' : 'Private'}
                </span>
              </h5>
              <p className="card-text">{project.description}</p>
              <p className="card-text">
                <span className="me-2">
                  <i className="fas fa-circle me-1"></i>
                  {project.language}
                </span>
                <i className="far fa-clock me-1"></i>
                {project.lastUpdated}
              </p>
            </div>
          </div>
        ))}
      </div>
    </>
  );
};

export default Repository;
