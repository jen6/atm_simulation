# ATM Simulation

## 설계
![image](https://github.com/user-attachments/assets/8000d00b-b33c-47bc-bd88-e3ed8eea58bd)

## 고려사항
- 은행별로 인증 방식이 다를 수 있기 때문에 session key를 account 마다 받게 한다
- account balance가 0원 이하거나 출금 금액보다 적어도 은행 Transaction이 허가된다면 출금한다 (마이너스 통장등, 인출 되는 것은 결국 은행의 결정사항)
- 입출금시 atm(cash bin)의 금액을 조정했다가 account에서 거래가 안된다면 해당 금액을 rollback한다.

## To-Do
- 입출금 처리는 됐지만 기계 고장으로 인해 입출금이 안됐다면 반대되는 거래를 다시 롤백한다
- atm 거래 관련 Transaction들을 별도로 저장한다
- unit test 작성 (현재 controller 테스트만 존재)

![image](https://github.com/user-attachments/assets/acc5ff06-3da1-4939-b78b-50e7daba769a)
![image](https://github.com/user-attachments/assets/b11a7fae-689a-4919-adbc-0a79fe8f02c1)
