import { redirect } from "@sveltejs/kit";

/** @type {import('@sveltejs/kit').PageData} */
export async function load({ parent, locals }) {
  const { session } = await parent();
  // or
  // locals.session.data.session;


  // Already logged in:
  if (session.user.accessToken) {
    throw redirect(302, '/')
  }

  return {};
}