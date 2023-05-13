require('./.env.validator');

/**
 * Don't be scared of the generics here.
 * All they do is to give us autocompletion when using this.
 *
 * @template {import('next').NextConfig} T
 * @param {T} config - A generic parameter that flows through to the return type
 * @constraint {{import('next').NextConfig}}
 */
function defineNextConfig(config) {
  return config;
}

module.exports = defineNextConfig({
  output: 'standalone',
  async redirects() {
    return [];
  },
  async rewrites() {
    return [
      // Rewrite everything else to use `pages/app`
      {
        source: '/app/:any*',
        destination: '/app/',
      },
    ];
  },
});
