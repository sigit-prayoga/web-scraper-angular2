import { WebScraperAngular2Page } from './app.po';

describe('web-scraper-angular2 App', function() {
  let page: WebScraperAngular2Page;

  beforeEach(() => {
    page = new WebScraperAngular2Page();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});
