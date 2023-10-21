import type { NextPage } from 'next';
import { Base } from '@/components/base/index';
import { Account } from '@/components/profile/index';
import { List } from '@/components/repository/index';
import styles from '../styles/Home.module.css';

const Home: NextPage = () => {
  return (
    <>
      <div className={`${styles.container} bg-light container`}>
        <Base.Header />

        <div className="mt-5">
          <div className="row">
            <div className='col-md-3'>
              <Account.Profile />
            </div>
            <div className='col-md-8'>
              <List.Repository />
            </div>
          </div>
        </div>

        <Base.Footer />
      </div>
    </>
  );
};

export default Home;
