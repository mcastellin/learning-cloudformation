from math import floor, ceil
from crhelper import CfnResource

helper = CfnResource()

@helper.create
@helper.update
def calculate_instances(event, _):
    if not 'DesiredInstances' in event['ResourceProperties']:
        helper.Data['Min'] = 0
        helper.Data['Max'] = 0
        helper.Data['Desired'] = 0
    else:
        desired = int(event['ResourceProperties']['DesiredInstances'])
        if desired == 0:
            helper.Data['Min'] = 0
            helper.Data['Max'] = 0
            helper.Data['Desired'] = 0
        else:
            min = floor(desired * 0.6)
            max = ceil(desired * 1.4)
            helper.Data['Min'] = min
            helper.Data['Max'] = max
            helper.Data['Desired'] = desired

@helper.delete
def no_op(_, __):
    pass

def lambda_handler(event, context):
    helper(event, context)