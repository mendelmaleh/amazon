# inflation

Attempt calculating the inflation rate by checking the difference of the items you purchased on Amazon in 2021 and their current price.

You can export your transaction history from the [Download order reports](https://www.amazon.com/gp/b2b/reports?ref_=ya_d_l_order_reports) page in your account settings, use the `Items` report type.


```sh
$ inflation -y 2021 -f items.csv
86 successful items, 30 failed items
2256.76 total 2021
2427.62 total now
+170.86 (+7%) difference
```
