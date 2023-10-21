// components/Header.js
import React from 'react';
import Link from 'next/link';
import { login } from '@/utils/auth';

const Header = () => {

  async function doLogin(e: any) {
    console.log("masuk")

    const d = await login()
    console.log(d)

  }

  return (
    <nav className='navbar navbar-light'>
      <div className='container-fluid'>
        <Link className='navbar-brand' href='#'>
          <img src='./main-logo.png' alt='' width='30' height='24' className='d-inline-block align-text-top' />
          <span>Simple.Repo</span>
        </Link>

        <div className='d-flex'>
            <button className='btn btn-primary' onClick={doLogin}>Login with Github</button>
        </div>
      </div>
    </nav>
  );
};

export default Header;
