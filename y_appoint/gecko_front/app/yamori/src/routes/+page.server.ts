/** @type {import('@sveltejs/kit').Actions} */
export const actions = {
  default: async ({ locals }) => {
    const { user: loginSession = {
      email: '',
      accessToken: ''
    } } = locals.session.data;

    await locals.session.set({
      user: {
        email: 'hoge@hoge.mail',
        accessToken: 'hogehogehogehogefugafugafugafuga'
      }
    });

    return {};
  }
};