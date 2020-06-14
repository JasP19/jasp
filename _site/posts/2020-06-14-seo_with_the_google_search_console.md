---
layout: layouts/post.njk
title: 'SEO with the Google Search Console'
summary: "Improving web application search performance using the Google Search Console."
hero: /images/posts/seo_with_the_google_search_console.png
thumb: /images/posts/seo_with_the_google_search_console.png
tags:
  - seo
  - google
  - web
---

## The Aim

It's fair to say that Google is the most-used search engine in the world. This is understandable considering how well-tuned its search algorithm is, meaning users get the most accurate results which are relevant to them. Software engineers are no strangers to the end-user experience of the Google search engine, however, there is often a different consideration that is equally as important. Developers who work on websites or web applications generally want their creations to appear as high up as possible in Google's search results. This greatly increases the chance that users will browse to the website and the process of enabling this is known as Search Engine Optimisation (SEO). Most software engineers know about SEO and understand its importance but many are not aware of the Google Search Console, one of the most powerful tools to make your web applications appear in Google's search results. The aim of this article is, then, to demonstrate how powerful the Google Search Console is and how it can be used to make your site appear in the results of the world's most popular search engine.

## How do we do it?

Before we get into the details let's take a look at what we need:
- A website or web application hosted on a public-facing domain
- A Google account
- Basic web-development knowledge

Before we can do anything else, we need to gain access to the Google Search Console. We do this by [navigating to the Console](https://search.google.com/search-console) and signing in using our Google account. We are then be presented with a welcome message and an important choice - whether we want to add our entire domain (with all URLs and prefixes) or whether we just want to add URLs with a certain prefix. In most cases, we will want to add our entire domain but advanced users may want to select only a subset of their domain's URLs to be indexed. Once we have made our decision, we enter our domain name and press 'Continue'. The next step is to verify that we actually own the domain in question. This can be done through a DNS record or a URL prefix property. In my opinion, the easiest is the DNS record option where we simply sign in to our account with our domain name registrar and add a TXT record that has the value provided by the Google Search Console (it will be a random string, prefixed with `google-site-verification`). Once ownership has been verified, we are redirected to the 'Overview' page of the Google Search Console.

Now that we have access to the search console, we move on to the most important step - ensuring all the pages (URLs) on our website are actually visible to Google's crawler so that they get indexed by the search engine. The easiest way to achieve this is through a Sitemap. A Sitemap is just an XML file which follows the [Sitemap protocol](https://www.sitemaps.org/protocol.html) and lists all of the different URLs which are accessible on our domain. The two most important properties are the `<loc>` and `<lastmod>` tags which represent the absolute URL and the date it was last modified, respectively. Once our Sitemap has been created, we deploy it to our site (preferably at the root level of the directory structure). We then select the 'Sitemaps' tab on the Google Search Console and submit the URL to our Sitemap. Google will verify that everything is in order (notifying us of any issues) and use the Sitemap to start indexing the URLs on our site. Let's take a look at an example of the sitemap for the site you're reading this article on (you can also [view the sitemap in your browser](https://jasp.dev/sitemap.xml)):

``` xml
<?xml version="1.0" encoding="utf-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">

  <url>
    <loc>https://jasp.dev/reviews/2017-03-25-horizon-zero-dawn-review/</loc>
    <lastmod>2017-03-25</lastmod>
  </url>
  
  <url>
    <loc>https://jasp.dev/posts/2020-05-02-what_does_a_software_engineer_do/</loc>
    <lastmod>2020-05-02</lastmod>
  </url>
  
  <url>
    <loc>https://jasp.dev/js/data.js</loc>
    <lastmod>2020-05-02</lastmod>
  </url>
  
  <url>
    <loc>https://jasp.dev/projects/</loc>
    <lastmod>2020-05-02</lastmod>
  </url>
  
  <url>
    <loc>https://jasp.dev/tags/</loc>
    <lastmod>2020-05-02</lastmod>
  </url>
  
  <url>
    <loc>https://jasp.dev/projects/IntercomPi/02_Requirements/</loc>
    <lastmod>2020-05-02</lastmod>
  </url>
  
  <url>
    <loc>https://jasp.dev/projects/IntercomPi/</loc>
    <lastmod>2020-05-02</lastmod>
  </url>
  
  <url>
    <loc>https://jasp.dev/reviews/0/</loc>
    <lastmod>2020-05-03</lastmod>
  </url>
  
  <url>
    <loc>https://jasp.dev/posts/0/</loc>
    <lastmod>2020-05-03</lastmod>
  </url>
  
  <url>
    <loc>https://jasp.dev/posts/2020-05-09-who_can_be_a_software_engineer/</loc>
    <lastmod>2020-05-09</lastmod>
  </url>
  
  <url>
    <loc>https://jasp.dev/posts/2020-05-16-email_with_firebase_functions_and_firestore/</loc>
    <lastmod>2020-05-16</lastmod>
  </url>
  
  <url>
    <loc>https://jasp.dev/reviews/2020-05-23-journey-to-the-savage-planet-review/</loc>
    <lastmod>2020-05-23</lastmod>
  </url>
  
  <url>
    <loc>https://jasp.dev/posts/2020-05-31-what_qualifications_do_software_engineers_need/</loc>
    <lastmod>2020-05-31</lastmod>
  </url>
  
  <url>
    <loc>https://jasp.dev/</loc>
    <lastmod>2020-06-06</lastmod>
  </url>
  
  <url>
    <loc>https://jasp.dev/about/</loc>
    <lastmod>2020-06-06</lastmod>
  </url>
  
  <url>
    <loc>https://jasp.dev/reviews/2020-06-07-valorant-a-valiant-idea/</loc>
    <lastmod>2020-06-07</lastmod>
  </url>
  
  <url>
    <loc>https://jasp.dev/posts/2020-06-14-seo_with_the_google_search_console/</loc>
    <lastmod>2020-06-14</lastmod>
  </url>
</urlset>
```

Once Google has used the Sitemap to index the URLs on our site - which takes up to 24 hours - we should start seeing our site in Google's search results, if we use the correct keywords that is. We are now also able to access some of the other features of the Google Search Console. The most important of these is the 'Performance' tab. As data builds up for our site, we will see metrics such as the total number of clicks on our site from users of the Google search engine, the number of times our site has appeared in search results, and the average position of our site in search results. There are a variety of filters for these metrics which allow us to monitor the performance of specific pages as well as details about the devices which are used to access our site and where (in the world) our site is accessed from. By focusing on these statistics in detail, we can ensure our pages appear higher in search results. This is done by improving various on-page SEO factors such as page speed or the number of relevant internal page links. Overall, this allows us to increase our click-through rate (CTR - the number of users who click on our site after seeing it in search results). Additionally, we will be informed of any issues on any of our pages (clickable elements being positioned too closely together on mobile devices, for example). It's worth fixing issues that crop up as soon as possible since they can negatively affect the chance of your site appearing in results.

Of course, the Google Search Console is not the only way to improve the SEO of your site. There are other important aspects such as adding meta tags to your pages and just increasing the amount of traffic to your site (through advertising, for example). That being said, the Google Search Console is an immensely powerful tool which I hope you'll use in future to improve the search ratings of your web applications.

_Jason Parry_
