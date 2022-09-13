// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
// and what to do when importing types
/// <reference types="@sveltejs/kit" />
declare namespace App {
	interface Locals extends UserSession {
		user: User;
		accessToken?: string;
		error: ApiError;
	}

	interface Session {
		user: User;
		accessToken?: string;
	}

	interface PageData {
		session: App.Session;
	}
	// interface Platform {}
}
