import { expect, test } from '@playwright/test';

test('index page has expected h1', async ({ page }) => {
	await page.goto('/');
	await expect(page.getByRole('heading', { name: 'Welcome to My App!' })).toBeVisible();
});

test('index page title', async ({page}) => {
  await page.goto('/');
  expect(await page.title() === "SPA");;
})
