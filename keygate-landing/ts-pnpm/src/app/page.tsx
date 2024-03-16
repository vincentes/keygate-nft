'use client';

import Head from 'next/head';
import * as React from 'react';
import '@/lib/env';
import { Button } from 'antd';
import { Card, Space } from 'antd';

/**
 * SVGR Support
 * Caveat: No React Props Type.
 *
 * You can override the next-env if the type is important to you
 * @see https://stackoverflow.com/questions/68103844/how-to-override-next-js-svg-module-declaration
 */
import Logo from '~/svg/Logo.svg';
import WalletExample from '~/svg/wallet.svg';

// !STARTERCONF -> Select !STARTERCONF and CMD + SHIFT + F
// Before you begin editing, follow all comments with `STARTERCONF`,
// to customize the default configuration.

export default function HomePage() {
  return (
    <main>
      <Head>
        <title>Keygate</title>
      </Head>
      <section className='bg-white'>
        <div className='layout mx-auto flex flex-col pt-20 items-center text-center'>
          <div className='mt-4 text-lg font-semibold text-gray-800 tracking-widest'>
            WEB3 UX & INFRASTRUCTURE PLATFORM
          </div>
          <h1 className='w-[600px] text-5xl font-semibold tracking-wide'>
            We onramp enterprises into scalable and safe blockchain
            architecture.
          </h1>
          <div className='w-[600px] text-xl'>
            With Keygate, seamlessly onboard your users to the world of
            blockchain through our intuitive APIs and no-code solutions,
            catering to both savvy and novice crypto users alike.
          </div>
          <div className='w-[600px] mt-20'>
            <Button type='text' size='large'>
              <a href='mailto:vicente@keygate.club'>Contact us</a>
            </Button>
          </div>
        </div>
        <section className=' bg-black h-[10px] pl-12 pt-40 pb-20 text-4xl tracking-widest text-center text-white mt-40'></section>
        <section className=' bg-white h-[10px] pl-12 pt-40 pb-20 text-4xl tracking-widest text-center text-white mt-20'></section>
      </section>
    </main>
  );
}
