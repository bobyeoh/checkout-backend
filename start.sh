docker run -it --name checkout-backend \
-p 0.0.0.0:8888:8888 \
-e ENV="dev" \
-e RDS="root:8506623@tcp(127.0.0.1:3306)/checkout?charset=utf8&parseTime=True&loc=Local" \
-d checkout-backend