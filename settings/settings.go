package settings

import (
	"os"
	"net/http"
	"fmt"
	"strconv"
)

var (
	AUTHOR = os.Getenv("AUTHOR")
	COMPANY_NAME = os.Getenv("COMPANY_NAME")
	COMPANY_NAME_ALTERNATE = os.Getenv("COMPANY_NAME_ALTERNATE")
	COMPANY_MAILING_ADDRESS = os.Getenv("COMPANY_MAILING_ADDRESS")
	COMPANY_CONTACT_EMAIL = os.Getenv("COMPANY_CONTACT_EMAIL")
	COMPANY_CONTACT_PHONE = os.Getenv("COMPANY_CONTACT_PHONE")
	COMPANY_SUPPORT_EMAIL = os.Getenv("COMPANY_SUPPORT_EMAIL")
	COMPANY_ORDERS_EMAIL = os.Getenv("COMPANY_ORDERS_EMAIL")
	TAX_PERCENT float64 = 0.0
	COMPANY_URL = os.Getenv("COMPANY_URL")
	COMPANY_GOOGLE_MAPS_URL = os.Getenv("COMPANY_GOOGLE_MAPS_URL")
	COMPANY_GOOGLE_MAPS_EMBED_URL = os.Getenv("COMPANY_GOOGLE_MAPS_EMBED_URL")
	FACEBOOK_URL = os.Getenv("FACEBOOK_URL")
	FACEBOOK_APP_ID = os.Getenv("FACEBOOK_APP_ID")
	INSTAGRAM_URL = os.Getenv("INSTAGRAM_URL")
	TWITTER_URL = os.Getenv("TWITTER_URL")
	LINKEDIN_URL = os.Getenv("LINKEDIN_URL")
	YOUTUBE_URL = os.Getenv("YOUTUBE_URL")
	TWITTER_HANDLE = os.Getenv("TWITTER_HANDLE")
	GOOGLE_ANALYTICS_ACCOUNT_ID = os.Getenv("GOOGLE_ANALYTICS_ACCOUNT_ID")
	GOOGLE_TAG_MANAGER_ID = os.Getenv("GOOGLE_TAG_MANAGER_ID")

	PAYPAL_ENVIRONMENT = os.Getenv("PAYPAL_ENVIRONMENT")
	PAYPAL_API_URL = os.Getenv("PAYPAL_API_URL")
	PAYPAL_EMAIL = os.Getenv("PAYPAL_EMAIL")
	PAYPAL_ACCOUNT = os.Getenv("PAYPAL_ACCOUNT")
	PAYPAL_API_CLIENT_ID = os.Getenv("PAYPAL_API_CLIENT_ID")
	PAYPAL_API_CLIENT_SECRET = os.Getenv("PAYPAL_API_CLIENT_SECRET")
	PAYPAL_ALLOWED_PAYMENT_OPTION = os.Getenv("PAYPAL_ALLOWED_PAYMENT_OPTION") //posible: UNRESTRICTED, INSTANT_FUNDING_SOURCE, IMMEDIATE_PAY
	PAYPAL_NOTE_TO_PAYER = os.Getenv("PAYPAL_NOTE_TO_PAYER")

	SMARTYSTREETS_AUTH_ID = os.Getenv("SMARTYSTREETS_AUTH_ID")
	SMARTYSTREETS_AUTH_TOKEN = os.Getenv("SMARTYSTREETS_AUTH_TOKEN")

	EMAIL_SENDER = os.Getenv("EMAIL_SENDER")
	SENDGRID_KEY = os.Getenv("SENDGRID_KEY")

	META_TITLE_HOME = os.Getenv("META_TITLE_HOME")

	META_DESCRIPTION_HOME = os.Getenv("META_DESCRIPTION_HOME")
	META_DESCRIPTION_STORE = os.Getenv("META_DESCRIPTION_STORE")
	META_DESCRIPTION_REFERRALS = os.Getenv("META_DESCRIPTION_REFERRALS")
	META_DESCRIPTION_CONTACT = os.Getenv("META_DESCRIPTION_CONTACT")
	META_DESCRIPTION_CART = os.Getenv("META_DESCRIPTION_CART")
	META_DESCRIPTION_BLOG = os.Getenv("META_DESCRIPTION_BLOG")
	META_DESCRIPTION_GALLERIES = os.Getenv("META_DESCRIPTION_GALLERIES")

	DESCRIPTION_BLOG_ABOUT = os.Getenv("DESCRIPTION_BLOG_ABOUT")
	WWW_REDIRECT bool
)

func init() {
	TAX_PERCENT, _ = strconv.ParseFloat(os.Getenv("TAX_PERCENT"), 64)
	WWW_REDIRECT, _ = strconv.ParseBool(os.Getenv("WWW_REDIRECT"))
}

func ServerUrl(r *http.Request) string {
	httpHeader := "http"
	if r.TLS != nil {
		httpHeader = "https"
	}

	return fmt.Sprintf("%s://%s", httpHeader, r.Host)
}