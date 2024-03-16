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
          <div className='mt-4 text-lg font-semibold text-primary-500 tracking-widest'>
            WEB3 UX & INFRASTRUCTURE PLATFORM
          </div>
          <h1 className='w-[600px] text-5xl font-semibold tracking-wide'>
            We onramp enterprises into scalable and safe blockchain
            architecture.
          </h1>
          <div className='w-[600px] text-xl'>
            With Supahub collect, organize and prioritize feature requests to
            better understand customer feedback and use them to inform your
            product roadmap.
          </div>
          <div className='w-[600px] mt-10'>
            <Button type='primary' size='large'>
              {' '}
              Sign up for free
            </Button>
          </div>
        </div>
        <section className='w-full bg-primary-500 pl-12 py-20 mt-20'>
          <div className='tracking-widest font-semibold'>
            GENERAL USE CASE PLATFORM
          </div>
          <div className='text-4xl w-[600px] font-semibold mt-5'>
            We give you the building blocks, all batteries included.
          </div>
          <div className='mt-10 rounded-xl pl-10 pt-10 pb-5 w-[600px]'>
            <WalletExample className='w-[200px]' />
          </div>
        </section>
      </section>
    </main>
  );
}
