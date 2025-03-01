# coding: utf-8

"""


    Generated by: https://openapi-generator.tech
"""

import unittest
from unittest.mock import patch

import urllib3

import bitget
from bitget.paths.api_margin_v1_cross_interest_list import get  # noqa: E501
from bitget import configuration, schemas, api_client

from .. import ApiTestMixin


class TestApiMarginV1CrossInterestList(ApiTestMixin, unittest.TestCase):
    """
    ApiMarginV1CrossInterestList unit test stubs
        list  # noqa: E501
    """
    _configuration = configuration.Configuration()

    def setUp(self):
        used_api_client = api_client.ApiClient(configuration=self._configuration)
        self.api = get.ApiForget(api_client=used_api_client)  # noqa: E501

    def tearDown(self):
        pass

    response_status = 200

    def testApi(self):
        configuration = ApiTestMixin.get_default_configuration()
        # Enter a context with an instance of the API client
        with bitget.ApiClient(configuration) as api_client:
            # Create an instance of the API class
            from bitget.apis.tags import margin_cross_interest_api
            api_instance = margin_cross_interest_api.MarginCrossInterestApi(api_client)
            try:
                req = {
                    'startTime': "1679133422000",
                }
                api_response = api_instance.cross_interest_list(req)
                print(api_response)
                self.assertIsNotNone(api_response)
                self.assertIsNotNone(api_response.body)
                self.assertEqual(api_response.body['code'], '00000')
                self.assertEqual(api_response.body['msg'], 'success')
                self.assertIsNotNone(api_response.body['data']['resultList'])
                for item in api_response.body['data']['resultList']:
                    print(item)
                    self.assertIsNotNone(item['interestId'])
                    self.assertIsNotNone(item['loanCoin'])
                    self.assertIsNotNone(item['amount'])
                    self.assertIsNotNone(item['type'])
            except bitget.ApiException as e:
                print("Exception when calling place_order: %s\n" % e)


if __name__ == '__main__':
    unittest.main()
