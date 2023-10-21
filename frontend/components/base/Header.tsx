// components/Header.js
import Link from 'next/link';

const Header = () => {
  return (
    <nav className='navbar navbar-light'>
      <div className='container-fluid'>
        <Link className='navbar-brand' href='#'>
          <img src='./main-logo.png' alt='' width='30' height='24' className='d-inline-block align-text-top' />
          <span>Simple.Repo</span>
        </Link>

        <form className='d-flex'>
          <Link href={''}>
            <button className='btn btn-primary'>Login with Github</button>
          </Link>
        </form>
      </div>
    </nav>
  );
};

export default Header;
