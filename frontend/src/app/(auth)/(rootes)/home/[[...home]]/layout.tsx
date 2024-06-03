import type { Metadata } from 'next'
import NextTopLoader from 'nextjs-toploader';

export const metadata: Metadata = {
    title: 'Home Page | ShortyURL',
    description: 'InCoder Web',
}

const SignInLayout = ({
  children,
}: {
  children: React.ReactNode
}) => {
  return (
    <>
      <NextTopLoader color="#000" showSpinner={false} />
      {children}
    </>
  );
};

export default SignInLayout;