func fix_image(url) {
  url = strings.Replace(url, "-sign-", "-", -1)
  url = strings.Replace(url, "-sign.", ".", -1)
  return strings.Split(url, "?")[0]
}

expired_image := "https://p16-sign-va.tiktokcdn.com/tos-maliva-avt-0068/9435f286c26b12a87f0ea4ee4adf07f9~c5_1080x1080.webp?x-expires=1661893200&x-signature=6nylSljOgwMGkxIiv%2BrEfIz0VoU%3D"

new_working_image := fix_image(expired_image)

fmt.Println(new_working_image)
